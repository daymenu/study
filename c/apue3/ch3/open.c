#include <fcntl.h>
#include <apue.h>

char buf1[] = "abcdefghij";
char buf2[] = "ABCDEFGHIJ";
int main(int argc, char *argv[])
{
    int fd;

    if ((fd = creat("file.hole", FILE_MODE)) < 0)
    {
        err_sys("creat error");
    }

    if (write(fd, buf1, 10) != 10)
    {
        err_sys("buf1 wirte error");
    }

    if (lseek(fd, 16384, SEEK_SET) == -1)
    {
        err_sys("lseek error");
    }
    
    if (write(fd, buf2, 10) != 10)
    {
        err_sys("buf2 wirte error");
    }
    return 0;
}