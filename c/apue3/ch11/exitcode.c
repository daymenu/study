#include <stdio.h>
#include <pthread.h>
#include <stdlib.h>

void *thr_fn1(void *arg)
{
    printf("thread 1 returing\n");
    return ((void *)1);
}

void *thr_fn2(void *arg)
{
    printf("thread 2 returing\n");
    return ((void *)2);
}

int main(int argc, char const *argv[])
{
    int err;
    pthread_t tid1, tid2;
    void *tret;

    err = pthread_create(&tid1, NULL, thr_fn1, NULL);
    if (0 != err)
    {
        printf("can't create thread 1\n");
        exit(err);
    }

    err = pthread_create(&tid2, NULL, thr_fn2, NULL);
    if (0 != err)
    {
        printf("can't create thread 2\n");
        exit(err);
    }

    err = pthread_join(tid1, &tret);
    if (err != 0)
    {
        printf("can't join with thread 1\n");
        exit(err);
    }
    printf("thread 1 exit coee %ld\n", (long)tret);

    err = pthread_join(tid2, &tret);
    if (err != 0)
    {
        printf("can't join with thread 2\n");
        exit(err);
    }
    printf("thread 2 exit coee %ld\n", (long)tret);
    return 0;
}
