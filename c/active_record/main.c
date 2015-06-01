#include <stdio.h>

#include "active_record.h"

int
main(int argc, char* argv[])
{
    session *sess = new_session();
    printf("%s\n", sess->select(sess, "id,name")->where(sess, "id = 1")->get(sess, "user")->query(sess));
    return 0;
}
