#include <iostream>
#include <string>
#include <cassert>
#include "utility.hpp"

using namespace std;

const int SCREEN_HEIGHT = 6;
const int SCREEN_WIDTH = 40;

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
	measure_time.push_back(260);

	int time = 1;
	auto measure_time_it = measure_time.begin();
	int x_val = 1;


	vector<vector<char>> screen(SCREEN_HEIGHT, vector<char>(SCREEN_WIDTH, '.'));

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

		for (int j = time; j < new_time; j++) {
			int pixel_y = (j - 1) / SCREEN_WIDTH;
			int pixel_x = (j - 1) % SCREEN_WIDTH;

			if (pixel_y <= SCREEN_HEIGHT - 1) {
				if (pixel_x >= x_val - 1 && pixel_x <= x_val + 1) {
					screen[pixel_y][pixel_x] = '#';
				}
				else {
					screen[pixel_y][pixel_x] = '.';
				}
			}
		}
		x_val = new_x_val;
		time = new_time;
	}

	{
		int j = time;
		int pixel_y = (j - 1) / SCREEN_WIDTH;
		int pixel_x = (j - 1) % SCREEN_WIDTH;

		if (pixel_y <= SCREEN_HEIGHT - 1) {
			if (pixel_x >= x_val - 1 && pixel_x <= x_val + 1) {
				screen[pixel_y][pixel_x] = '#';
			}
			else {
				screen[pixel_y][pixel_x] = '.';
			}
		}
	}



	int result = 0;
	for (auto signal_strength : signal_strengths) {
		result += signal_strength;
	}
	cout << result << "\n";


	for (int i = 0; i < SCREEN_HEIGHT; i++) {
		for (int j = 0; j < SCREEN_WIDTH; j++) {
			cout << screen[i][j];
		}
		cout << "\n";
	}

	return 0;
}
