#include <iostream>
#include <string>
#include <cassert>
#include "utility.hpp"

using namespace std;

int main() {

	vector<string> instructions;
	string line;
	while (getline(cin, line)) {
		instructions.push_back(line);
	}

	vector<int> measure_time;
	vector<int> signal_strengths;
	for (int i = 20; i <=220; i += 40) {
		measure_time.push_back(i);
	}

	int time = 1;
	auto measure_time_it = measure_time.begin();
	int x_val = 1;

	for (int i = 0; i < instructions.size() && measure_time_it != measure_time.end(); i++) {
		const auto& instruction = instructions[i];

		auto words = split(instruction);
		assert(!words.empty());
		auto new_x_val = x_val;
		auto new_time = time;
		if (words[0] == "noop") {
			new_time = time + 1;
		}
		else if (words[0] == "addx") {
			assert(words.size() >= 2);
			new_time = time + 2;
			new_x_val = x_val + stoi(words[1]);
		}

		if (new_time > *measure_time_it) {
			signal_strengths.push_back(*measure_time_it * x_val);
			measure_time_it++;
		}
		else if (new_time == *measure_time_it) {
			signal_strengths.push_back(*measure_time_it * new_x_val);
			measure_time_it++;
		}
		x_val = new_x_val;
		time = new_time;
	}


	int result = 0;
	for (auto signal_strength : signal_strengths) {
		result += signal_strength;
	}
	cout << result << "\n";


	return 0;
}
