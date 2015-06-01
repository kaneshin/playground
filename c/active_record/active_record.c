/**
 * active_record.h
 *
 * Copyright (c) 2015 Shintaro Kaneko
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in
 * all copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
 * THE SOFTWARE.
 */

#include "active_record.h"

#include <stdio.h>
#include <stdlib.h>

struct internal_session {
    char *table;
    char *cols;
    char *cond;
    char buf[256];
};
typedef struct internal_session internal_session;

session*
session_get(session *self, char *table)
{
    self->_internal->table = table;
    return self;
}

session*
session_select(session *self, char *cols)
{
    self->_internal->cols = cols;
    return self;
}

session*
session_where(session *self, char *cond)
{
    self->_internal->cond = cond;
    return self;
}

char*
session_query(session *self)
{
    sprintf(self->_internal->buf, "SELECT %s FROM %s WHERE %s", self->_internal->cols, self->_internal->table, self->_internal->cond);
    return self->_internal->buf;
}

session*
new_session()
{
    session *sess = (session *)malloc(sizeof(session));
    if (sess == NULL)
        return NULL;

    internal_session *internal = (internal_session *)malloc(sizeof(internal_session));
    if (internal == NULL)
    {
        free(sess);
        return NULL;
    }
    internal->table = NULL;
    internal->cols = NULL;
    internal->cond = NULL;
    sess->_internal = internal;
    sess->get = session_get;
    sess->select = session_select;
    sess->where = session_where;
    sess->query = session_query;
    return sess;
}

void
free_session(session *sess)
{
    if (sess != NULL)
    {
        if (sess->_internal != NULL)
            free(sess->_internal);

        free(sess);
    }
}
