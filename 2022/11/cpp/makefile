CC = g++
FLAGS = -I .
DEPS = utility.hpp
OBJS = monkeys.o utility.o


monkeys.out: $(OBJS)
	$(CC) -o $@ $(OBJS) $(FLAGS)

%.o: %.c $(DEPS)
	$(CC) -c -o $@ $< $(FLAGS)

clean:
	rm -f *.o
	rm -f *.out
