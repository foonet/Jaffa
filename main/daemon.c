// ZmEu (zmeu_at_whitehat.ro) 09092018
//  - detaches from terminal and runs in background with first option 1/0 if to not chdir("/") and second option
//    is 1/0 if to keep standard file handles open
// gcc -o daemon daemon.c && ./daemon 1 1 /usr/bin/id
// uid=0(root) gid=0(root) groups=0(root),1(bin),2(daemon),3(sys),4(adm),6(disk),10(wheel),11(floppy),26(tape),27(video)

#include <unistd.h>
#include <stdio.h>
#include <errno.h>
#include <stdlib.h>
#include <string.h>

int main (int argc, char **argv) {
                int error;
                if (argc < 4) {
                        fprintf(stderr, "[-] syntax: %s\n<%%i>\t\t-chdir to root or not (booleon)\n<%%i>\t\t-keep standard file handles open or not (boolean)\n<%%s>\t\t-program to daemonize\n[$
                        return 1;
                }

                if (daemon(atoi(argv[1]), atoi(argv[2])) == -1) {
                        fprintf(stderr, "[-] there was an error to daemon(): %s", strerror(errno));
                        return 1;
                }

                if (error = execv(argv[3], argv+3)) {
                        fprintf(stderr, "[-] execv() returned an error: %s\n", strerror(errno));
                }

                return error;
}
