#include <stdlib.h>
#include <signal.h>
#include <stdio.h>
#include <unistd.h>

void abort(void)
{
    sigset_t mask;
    struct sigaction action;
    // Caller can't ignore SIGABRT,if so reset to default
    sigaction(SIGABRT, NULL, &action);
    if (action.sa_handler == SIG_IGN)
    {
        action.sa_handler = SIG_DFL;
        sigaction(SIGABRT, &action, NULL);
    }

    if (action.sa_handler == SIG_DFL)
    {
        fflush(NULL); // flush all open stdio streams
    }

    // Caller can't block SIGABRT;make sure it's unblock
    sigfillset(&mask);
    sigdelset(&mask, SIGABRT); // mask has only SIGABRT turned off
    sigprocmask(SIG_SETMASK, &mask, NULL);
    kill(getpid(), SIGABRT);

    // If we're here process caught SIGABRT and returned
    fflush(NULL);
    action.sa_handler = SIG_DFL;
    sigaction(SIGABRT, &action, NULL);
    sigprocmask(SIG_SETMASK, &action, NULL);
    kill(getpid(), SIGABRT);
    exit(1); // this should never be executed....
}
