package services

import "testing"

func TestLifecycleResetFLRD3MSEBMETransitions(t *testing.T) {
	var lifecycle Lifecycle

	out := lifecycle.Apply(Inputs{DState: 0, MSE: true, BME: true})
	if out.DeviceReset || !out.IOEnabled || !out.DMAEnabled || out.Quiesce {
		t.Fatalf("D0 with MSE+BME should be fully enabled, got %+v", out)
	}
	initialGeneration := out.Generation

	out = lifecycle.Apply(Inputs{DState: 0, MSE: false, BME: true})
	if out.IOEnabled || !out.DMAEnabled {
		t.Fatalf("MSE must gate only MMIO, got %+v", out)
	}
	if out.Generation != initialGeneration {
		t.Fatalf("ordinary MSE gating reset state: generation %d -> %d", initialGeneration, out.Generation)
	}

	out = lifecycle.Apply(Inputs{DState: 0, MSE: true, BME: false})
	if !out.IOEnabled || out.DMAEnabled {
		t.Fatalf("BME must gate only DMA, got %+v", out)
	}
	if out.Generation != initialGeneration {
		t.Fatalf("ordinary BME gating reset state: generation %d -> %d", initialGeneration, out.Generation)
	}

	out = lifecycle.Apply(Inputs{DState: 3, MSE: true, BME: true})
	if out.IOEnabled || out.DMAEnabled || !out.Quiesce {
		t.Fatalf("D3 must quiesce both MMIO and DMA, got %+v", out)
	}
	if out.Generation != initialGeneration {
		t.Fatalf("D3 transition must retain function state: generation %d -> %d", initialGeneration, out.Generation)
	}

	out = lifecycle.Apply(Inputs{DState: 0, MSE: true, BME: true})
	if !out.IOEnabled || !out.DMAEnabled || out.Quiesce || out.Generation != initialGeneration {
		t.Fatalf("D0 resume should restore gates without resetting retained state, got %+v", out)
	}

	out = lifecycle.Apply(Inputs{Reset: true, DState: 0, MSE: true, BME: true})
	if !out.DeviceReset || out.IOEnabled || out.DMAEnabled || !out.Quiesce {
		t.Fatalf("reset must quiesce all activity, got %+v", out)
	}
	if out.Generation != initialGeneration+1 {
		t.Fatalf("reset rising edge generation = %d, want %d", out.Generation, initialGeneration+1)
	}
	resetGeneration := out.Generation

	out = lifecycle.Apply(Inputs{Reset: true, DState: 0, MSE: true, BME: true})
	if out.Generation != resetGeneration {
		t.Fatalf("held reset incremented generation again: got %d, want %d", out.Generation, resetGeneration)
	}
	lifecycle.Apply(Inputs{DState: 0, MSE: true, BME: true})

	out = lifecycle.Apply(Inputs{FLR: true, DState: 0, MSE: true, BME: true})
	if !out.DeviceReset || out.IOEnabled || out.DMAEnabled || !out.Quiesce {
		t.Fatalf("FLR must quiesce all activity, got %+v", out)
	}
	if out.Generation != resetGeneration+1 {
		t.Fatalf("FLR rising edge generation = %d, want %d", out.Generation, resetGeneration+1)
	}
	flrGeneration := out.Generation
	out = lifecycle.Apply(Inputs{FLR: true, DState: 0, MSE: true, BME: true})
	if out.Generation != flrGeneration {
		t.Fatalf("held FLR incremented generation again: got %d, want %d", out.Generation, flrGeneration)
	}
}

func TestLifecycleTurnoffAndLinkDownQuiesceWithoutReset(t *testing.T) {
	for _, tc := range []struct {
		name string
		in   Inputs
	}{
		{name: "turnoff", in: Inputs{DState: 0, MSE: true, BME: true, Turnoff: true}},
		{name: "link_down", in: Inputs{DState: 0, MSE: true, BME: true, LinkDown: true}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			var lifecycle Lifecycle
			before := lifecycle.Apply(Inputs{DState: 0, MSE: true, BME: true})
			after := lifecycle.Apply(tc.in)
			if !after.Quiesce || after.IOEnabled || after.DMAEnabled {
				t.Fatalf("transition should quiesce traffic, got %+v", after)
			}
			if after.DeviceReset || after.Generation != before.Generation {
				t.Fatalf("traffic quiesce must not reset retained state: before=%+v after=%+v", before, after)
			}
		})
	}
}

func TestTagAllocatorAllocationWrapCompletionAndError(t *testing.T) {
	allocator := NewTagAllocator(30, 4, 100)

	var tags []uint8
	for range 4 {
		tag, ok := allocator.Allocate(10)
		if !ok {
			t.Fatal("allocator exhausted before all configured tags were allocated")
		}
		tags = append(tags, tag)
	}
	want := []uint8{30, 31, 32, 33}
	for i := range want {
		if tags[i] != want[i] {
			t.Fatalf("tag[%d] = %d, want %d (all tags %v)", i, tags[i], want[i], tags)
		}
	}
	if tag, ok := allocator.Allocate(10); ok {
		t.Fatalf("exhausted allocator returned tag %d", tag)
	}

	completed, ok := allocator.Complete(tags[1])
	if !ok || completed.Tag != tags[1] || completed.Kind != Completed {
		t.Fatalf("completion mismatch: outcome=%+v ok=%v", completed, ok)
	}
	if _, ok := allocator.Complete(tags[1]); ok {
		t.Fatal("duplicate completion was accepted")
	}

	reused, ok := allocator.Allocate(20)
	if !ok {
		t.Fatal("completed tag was not returned to the free pool")
	}
	if reused != tags[1] {
		t.Fatalf("allocation cursor did not wrap to the freed tag: got %d, want %d", reused, tags[1])
	}
	failed, ok := allocator.Fail(reused)
	if !ok || failed.Tag != reused || failed.Kind != Error {
		t.Fatalf("error completion mismatch: outcome=%+v ok=%v", failed, ok)
	}
	if _, ok := allocator.Fail(reused); ok {
		t.Fatal("duplicate error completion was accepted")
	}
	if _, ok := allocator.Complete(200); ok {
		t.Fatal("completion for an unallocated tag was accepted")
	}
}

func TestTagAllocatorTimeoutCancelAndNoStaleActivity(t *testing.T) {
	allocator := NewTagAllocator(7, 3, 10)
	tag0, ok := allocator.Allocate(100)
	if !ok {
		t.Fatal("first allocation failed")
	}
	tag1, ok := allocator.Allocate(105)
	if !ok {
		t.Fatal("second allocation failed")
	}

	if outcomes := allocator.Tick(109); len(outcomes) != 0 {
		t.Fatalf("tag timed out early: %+v", outcomes)
	}
	outcomes := allocator.Tick(110)
	if len(outcomes) != 1 || outcomes[0].Tag != tag0 || outcomes[0].Kind != Timeout {
		t.Fatalf("deadline should time out only tag %d, got %+v", tag0, outcomes)
	}
	if _, ok := allocator.Complete(tag0); ok {
		t.Fatal("late completion after timeout was accepted")
	}
	if outcomes := allocator.Tick(114); len(outcomes) != 0 {
		t.Fatalf("second tag timed out early: %+v", outcomes)
	}

	cancelled := allocator.CancelAll()
	if len(cancelled) != 1 || cancelled[0].Tag != tag1 || cancelled[0].Kind != Cancelled {
		t.Fatalf("cancel should report the sole outstanding tag, got %+v", cancelled)
	}
	if again := allocator.CancelAll(); len(again) != 0 {
		t.Fatalf("repeated cancel produced stale outcomes: %+v", again)
	}
	if _, ok := allocator.Complete(tag1); ok {
		t.Fatal("late completion after reset/cancel was accepted")
	}
	if outcomes := allocator.Tick(1_000); len(outcomes) != 0 {
		t.Fatalf("cancelled tags produced stale timeout activity: %+v", outcomes)
	}

	for i := range 3 {
		if _, ok := allocator.Allocate(2_000); !ok {
			t.Fatalf("allocator did not recover all capacity after timeout/cancel at allocation %d", i)
		}
	}
}

func TestInterruptControllerMSIEnableMasking(t *testing.T) {
	interrupts := NewInterruptController(1)

	if delivery := interrupts.Request(0); delivery.Valid {
		t.Fatalf("disabled MSI delivered immediately: %+v", delivery)
	}
	if !interrupts.Pending(0) {
		t.Fatal("disabled MSI request was not retained pending")
	}
	deliveries := interrupts.SetEnabled(true)
	assertSingleDelivery(t, deliveries, 0)
	if interrupts.Pending(0) {
		t.Fatal("MSI pending bit remained set after delivery")
	}
	if deliveries := interrupts.SetEnabled(true); len(deliveries) != 0 {
		t.Fatalf("re-enabling MSI redelivered an already consumed request: %+v", deliveries)
	}

	if delivery := interrupts.Request(0); !delivery.Valid || delivery.Vector != 0 {
		t.Fatalf("enabled MSI did not deliver immediately: %+v", delivery)
	}
	if interrupts.Pending(0) {
		t.Fatal("immediately delivered MSI incorrectly set pending")
	}
}

func TestInterruptControllerMSIXMasksPBARetentionAndExactlyOnceUnmask(t *testing.T) {
	interrupts := NewInterruptController(4)
	if deliveries := interrupts.SetEnabled(true); len(deliveries) != 0 {
		t.Fatalf("enabling idle controller delivered interrupts: %+v", deliveries)
	}

	interrupts.SetVectorMask(2, true)
	if delivery := interrupts.Request(2); delivery.Valid {
		t.Fatalf("vector-masked request delivered immediately: %+v", delivery)
	}
	if delivery := interrupts.Request(2); delivery.Valid {
		t.Fatalf("duplicate masked request delivered immediately: %+v", delivery)
	}
	if !interrupts.Pending(2) {
		t.Fatal("masked vector did not set its pending bit")
	}
	deliveries := interrupts.SetVectorMask(2, false)
	assertSingleDelivery(t, deliveries, 2)
	if interrupts.Pending(2) {
		t.Fatal("pending bit was not cleared by unmask delivery")
	}
	if deliveries := interrupts.SetVectorMask(2, false); len(deliveries) != 0 {
		t.Fatalf("repeated vector unmask redelivered consumed request: %+v", deliveries)
	}

	interrupts.SetFunctionMask(true)
	interrupts.SetVectorMask(1, true)
	if delivery := interrupts.Request(1); delivery.Valid {
		t.Fatalf("function/vector-masked request delivered immediately: %+v", delivery)
	}
	if delivery := interrupts.Request(3); delivery.Valid {
		t.Fatalf("second function-masked request delivered immediately: %+v", delivery)
	}
	if deliveries := interrupts.SetVectorMask(1, false); len(deliveries) != 0 {
		t.Fatalf("vector unmask bypassed function mask: %+v", deliveries)
	}
	if !interrupts.Pending(1) || !interrupts.Pending(3) {
		t.Fatal("PBA was lost while the function mask remained asserted")
	}
	deliveries = interrupts.SetFunctionMask(false)
	assertSingleDelivery(t, deliveries, 1)
	if interrupts.Pending(1) {
		t.Fatal("PBA was not cleared after function-unmask delivery")
	}
	if !interrupts.Pending(3) {
		t.Fatal("function unmask dropped a second pending vector instead of retaining it for drain")
	}
	if delivery := interrupts.Drain(); !delivery.Valid || delivery.Vector != 3 {
		t.Fatalf("drain did not deliver the second pending vector exactly once: %+v", delivery)
	}
	if delivery := interrupts.Drain(); delivery.Valid {
		t.Fatalf("drain redelivered consumed pending activity: %+v", delivery)
	}
	if deliveries := interrupts.SetFunctionMask(false); len(deliveries) != 0 {
		t.Fatalf("repeated function unmask redelivered consumed request: %+v", deliveries)
	}
}

func TestInterruptControllerResetDropsPendingWithoutStaleDelivery(t *testing.T) {
	interrupts := NewInterruptController(3)
	interrupts.SetEnabled(true)
	interrupts.SetFunctionMask(true)
	interrupts.Request(0)
	interrupts.Request(2)
	if !interrupts.Pending(0) || !interrupts.Pending(2) {
		t.Fatal("test setup failed to create pending vectors")
	}

	interrupts.Reset()
	for vector := range 3 {
		if interrupts.Pending(vector) {
			t.Fatalf("reset retained stale pending vector %d", vector)
		}
	}
	if deliveries := interrupts.SetFunctionMask(false); len(deliveries) != 0 {
		t.Fatalf("unmask after reset emitted stale interrupts: %+v", deliveries)
	}
	if deliveries := interrupts.SetEnabled(true); len(deliveries) != 0 {
		t.Fatalf("re-enable after reset emitted stale interrupts: %+v", deliveries)
	}
}

func TestDeviceResetCancelsOutstandingServicesWithoutStaleActivity(t *testing.T) {
	var lifecycle Lifecycle
	tags := NewTagAllocator(8, 2, 10)
	interrupts := NewInterruptController(2)

	lifecycle.Apply(Inputs{DState: D0, MSE: true, BME: true})
	tag, ok := tags.Allocate(100)
	if !ok {
		t.Fatal("test setup could not allocate a DMA tag")
	}
	interrupts.SetEnabled(true)
	interrupts.SetFunctionMask(true)
	interrupts.Request(1)

	reset := lifecycle.Apply(Inputs{Reset: true, DState: D0, MSE: true, BME: true})
	if !reset.DeviceReset {
		t.Fatal("test setup did not enter device reset")
	}
	cancelled := tags.CancelAll()
	interrupts.Reset()
	if len(cancelled) != 1 || cancelled[0].Tag != tag || cancelled[0].Kind != Cancelled {
		t.Fatalf("device reset did not cancel the outstanding tag: %+v", cancelled)
	}

	if _, ok := tags.Complete(tag); ok {
		t.Fatal("pre-reset completion was accepted after device reset")
	}
	if outcomes := tags.Tick(1_000); len(outcomes) != 0 {
		t.Fatalf("pre-reset tag generated stale timeout activity: %+v", outcomes)
	}
	if deliveries := interrupts.SetFunctionMask(false); len(deliveries) != 0 {
		t.Fatalf("pre-reset interrupt escaped after unmask: %+v", deliveries)
	}
	if deliveries := interrupts.SetEnabled(true); len(deliveries) != 0 {
		t.Fatalf("pre-reset interrupt escaped after re-enable: %+v", deliveries)
	}
}

func assertSingleDelivery(t *testing.T, deliveries []Delivery, vector int) {
	t.Helper()
	if len(deliveries) != 1 || !deliveries[0].Valid || deliveries[0].Vector != vector {
		t.Fatalf("deliveries = %+v, want one valid delivery for vector %d", deliveries, vector)
	}
}
