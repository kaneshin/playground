#include <stdio.h>
#include <stdlib.h>

typedef struct node_t
{
    struct node_t *prev;
    int data;
} node;

node*
new_node(int data)
{
    node *n = (node *)malloc(sizeof(node));
    n->data = data;
    return n;
}

void
free_node(node *n)
{
    if (NULL != n)
    {
        free(n);
        n = NULL;
    }
}

node*
append(node *prev, int data)
{
    node *next = new_node(data);
    next->prev = prev;
    return next;
}

void
delete_if(node *p, int data)
{
    node *t = NULL;
    while (p)
    {
        if (p->data != data)
        {
            p = p->prev;
            continue;
        }


    }
}

int
main(int argc, char* argv[])
{
    node *p = new_node(0);
    p = append(p, 1);
    p = append(p, 3);
    p = append(p, 3);
    p = append(p, 3);
    p = append(p, 3);
    p = append(p, 2);

    while (p)
    {
        printf("%d\n", p->data);
        p = p->prev;
    }

    printf("===\n");

    delete_if(p, 3);

    while (p)
    {
        printf("%d\n", p->data);
        p = p->prev;
    }


    return 0;
}
