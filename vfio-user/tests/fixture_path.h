#ifndef PCILEECH_VFIO_TEST_FIXTURE_PATH_H
#define PCILEECH_VFIO_TEST_FIXTURE_PATH_H

#include <limits.h>
#include <stdio.h>
#include <stdlib.h>

/*
 * Unit tests must consume fixtures generated into vfio-user/build rather than
 * stale Cocotb output directories.  The Makefile provides this variable for
 * every fixture-dependent test target; do not add a fallback path here.
 */
static inline int vfio_test_fixture_path(char output[PATH_MAX], const char *name)
{
    const char *root = getenv("VFIO_TEST_FIXTURE_ROOT");
    int written;

    if (root == NULL || root[0] == '\0' || name == NULL || name[0] == '\0')
        return -1;
    written = snprintf(output, PATH_MAX, "%s/%s", root, name);
    return written >= 0 && written < PATH_MAX ? 0 : -1;
}

#endif
