CC = g++
FLAGS = -I .
DEPS = utility.hpp
OBJS = crt.o utility.o


crt.out: $(OBJS)
	$(CC) -o $@ $(OBJS) $(FLAGS)

%.o: %.c $(DEPS)
	$(CC) -c -o $@ $< $(FLAGS)

clean:
	rm -f *.o
	rm -f *.out
