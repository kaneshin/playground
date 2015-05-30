// File:    person.c
// Author:  kaneshin <kaneshin0120@gmail.com>

#include "person.h"

#include <stdio.h>
#include <stdlib.h>

void person_set_name(person *, char *);
void person_set_gender(person *, gender);
void person_print(person *);

char *
gender_string(gender gen)
{
    switch (gen) {
    case GENDER_NEUTRAL:
        return "Neutral";
    case GENDER_MALE:
        return "Male";
    case GENDER_FEMALE:
        return "Female";
    case GENDER_UNKNOWN:
        return "Unknown";
    }
}

struct _internal_person {
    char *name;
    gender gen;
};

internal_person*
new_internal_person()
{
    internal_person *p = (internal_person *)malloc(sizeof(internal_person));
    if (p == NULL)
        return NULL;

    p->name = NULL;
    p ->gen = GENDER_UNKNOWN;
    return p;
}

void
free_internal_person(internal_person *p)
{
    if (NULL != p)
        free(p);
}

person*
new_person()
{
    person *p = (person *)malloc(sizeof(person));
    if (p == NULL)
        return NULL;

    internal_person *inter = new_internal_person();
    if (inter == NULL)
        return NULL;

    p->_internal = inter;
    p->set_name = person_set_name;
    p->set_gender = person_set_gender;
    p->print = person_print;
    return p;
}

void
free_person(person *p)
{
    if (NULL != p->_internal)
        free_internal_person(p->_internal);

    if (NULL != p)
        free(p);
}

void
person_set_name(person *self, char *name)
{
    self->_internal->name = name;
}

void
person_set_gender(person *self, gender gen)
{
    self->_internal->gen = gen;
}

void
person_print(person *self)
{
    printf("%s %s\n", self->_internal->name, gender_string(self->_internal->gen));
}

