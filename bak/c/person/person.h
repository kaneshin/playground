// File:    person.h
// Author:  kaneshin <kaneshin0120@gmail.com>

// gender

enum _gender
{
    GENDER_NEUTRAL,
    GENDER_MALE,
    GENDER_FEMALE,
    GENDER_UNKNOWN,
};
typedef enum _gender gender;

// person

typedef struct _internal_person internal_person;

struct _person {
    internal_person *_internal;

    void (*set_name)(struct _person *, char *name);
    void (*set_gender)(struct _person *, gender gen);
    void (*print)(struct _person *);
};
typedef struct _person person;

person* new_person();
void free_person(person *);
