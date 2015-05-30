// File:    main.c
// Author:  kaneshin <kaneshin0120@gmail.com>

#include "person.h"

int
main(int argc, char* argv[])
{
    person *p = new_person();
    p->set_name(p, "kaneshin");
    p->set_gender(p, GENDER_MALE);
    p->print(p);
    free_person(p);
    return 0;
}
