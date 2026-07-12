#include <signal.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#include "device_behavior.h"
#include "device_model.h"
#include "vfio_device.h"


static volatile sig_atomic_t stop_requested;


static void request_stop(int signal_number)
{
    (void)signal_number;
    stop_requested = 1;
}


static int usage(const char *program)
{
    fprintf(stderr, "usage: %s --artifacts DIR --socket PATH\n", program);
    return 2;
}


int main(int argc, char **argv)
{
    const char *artifacts = NULL;
    const char *socket_path = NULL;
    struct device_model *model = NULL;
    struct device_behavior behavior = {0};
    struct sigaction action = {0};
    char err[256] = {0};
    int index;
    int result = 1;

    for (index = 1; index < argc; ++index) {
        if (strcmp(argv[index], "--artifacts") == 0 && index + 1 < argc) {
            artifacts = argv[++index];
        } else if (strcmp(argv[index], "--socket") == 0 && index + 1 < argc) {
            socket_path = argv[++index];
        } else {
            return usage(argv[0]);
        }
    }
    if (artifacts == NULL || socket_path == NULL) {
        return usage(argv[0]);
    }
    if (device_model_load(artifacts, &model, err, sizeof(err)) < 0) {
        fprintf(stderr, "device model: %s\n", err);
        goto done;
    }
    if (behavior_create(model, &behavior, err, sizeof(err)) < 0) {
        fprintf(stderr, "device behavior: %s\n", err);
        goto done;
    }
    action.sa_handler = request_stop;
    sigemptyset(&action.sa_mask);
    if (sigaction(SIGINT, &action, NULL) < 0 || sigaction(SIGTERM, &action, NULL) < 0) {
        perror("sigaction");
        goto done;
    }
    if (vfio_device_run(model, &behavior, socket_path, &stop_requested) < 0) {
        perror("vfio device");
        goto done;
    }
    result = 0;

done:
    if (behavior.destroy != NULL) {
        behavior.destroy(behavior.state);
    }
    device_model_free(model);
    return result;
}
