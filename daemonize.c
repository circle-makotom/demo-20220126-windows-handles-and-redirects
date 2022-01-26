#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>

int main(void)
{
    pid_t pid;

    if ((pid = fork()) == 0)
    {
        printf("forked, daemonizing\n");

        close(0);
        close(1);
        close(2);

        sleep(30);
    }
    else if (pid > 0)
    {
        printf("forking\n");
    }

    return 0;
}
