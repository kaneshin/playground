SRCS = \
	main.c \
	person.c

OBJS = $(subst .c,.o,$(SRCS))

CFLAGS = 
LIBS = 
TARGET = a.out

all : $(TARGET)

$(TARGET) : $(OBJS)
	gcc -o $@ $(OBJS) $(LIBS)

.c.o :
	gcc -c $(CFLAGS) -I. $< -o $@

clean :
	rm -f *.o $(TARGET)
