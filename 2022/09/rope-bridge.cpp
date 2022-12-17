#include <string>
#include <iostream>
#include <vector>
#include <set>
#include <sstream>
#include <utility>
#include <cassert>



using namespace std;

pair<int, int> operator+(const pair<int, int> &a, const pair<int, int> &b);
pair<int, int> operator-(const pair<int, int> &a, const pair<int, int> &b);
pair<int, int> operator/(const pair<int, int> &a, int div);
pair<int, int> get_delta(string d);
int	max_abs_reduce(pair<int, int> p);
int get_manhattan_distance(pair<int, int> a, pair<int, int> b);
pair<int, int> get_new_position(pair<int, int> cur, pair<int, int> master_prev, pair<int, int> master_cur);
vector<string> split(const string& str);

void pp(const string &prefix, const pair<int, int> &p); //for debugging purposes

const int KNOTS_CNT = 10;

int main() {
	string line;

	set<pair<int, int>> visited_positions;
	visited_positions.insert(pair<int, int>(0, 0));
	vector<pair<int, int>> knots_pos(KNOTS_CNT, pair<int, int>(0, 0));

	while (getline(cin, line)) {
		auto words = split(line);
		string direction = words[0];
		int moves = stoi(words[1]);

		pair<int, int> delta = get_delta(direction);
		for (int i = 0; i < moves; i++) {
			auto &head_pos = knots_pos[0];
			auto master_prev_pos = head_pos;
			head_pos = head_pos + delta;
			for (int j = 1; j < KNOTS_CNT; j++) {
				//auto head_pos_prev = head_pos;
				const auto &master_cur_pos = knots_pos[j - 1];
				auto &cur_pos = knots_pos[j];


				{
					assert(max_abs_reduce(master_cur_pos - master_prev_pos) <= 1);
				}



				auto new_pos = get_new_position(cur_pos, master_prev_pos, master_cur_pos);
				master_prev_pos = cur_pos;
				cur_pos = new_pos;
			}
			const auto &tail_pos = knots_pos.back();
			visited_positions.insert(tail_pos);
		}
	}

	cout << visited_positions.size() << "\n";
	// for (auto p : visited_positions) {
	// 	pp("", p);
	// }
	return 0;
}

pair<int, int> operator+(const pair<int, int> &a, const pair<int, int> &b) {
	return pair<int, int>(a.first + b.first, a.second + b.second);
}

pair<int, int> operator-(const pair<int, int> &a, const pair<int, int> &b) {
	return pair<int, int>(a.first - b.first, a.second - b.second);
}

pair<int, int> operator/(const pair<int, int> &a, int div) {
	return pair<int, int>(a.first / div, a.second / div);
}

int max_abs_reduce(pair<int, int> p) {
	return max(abs(p.first), abs(p.second));
}

pair<int, int> get_delta(string d) {
	pair<int, int> res;
	
	if (d == "R") {
		res = pair<int, int>(0, 1);
	}
	else if (d == "L") {
		res = pair<int, int>(0, -1);
	}
	else if (d == "U") {
		res = pair<int, int>(1, 0);
	}
	else if (d == "D") {
		res = pair<int, int>(-1, 0);
	}
	return res;
}


int get_manhattan_distance(pair<int, int> a, pair<int, int> b) {
	return abs(a.first - b.first) + abs(a.second - b.second);
}

pair<int, int> get_new_position(pair<int, int> cur, pair<int, int> master_prev, pair<int, int> master_cur) {
	if (max_abs_reduce(cur - master_cur) <= 1) {
		return cur;
	}
	if (get_manhattan_distance(master_prev, master_cur) <= 1) {
		return master_prev;
	}


	auto md = get_manhattan_distance(cur, master_cur);
	assert(md >= 2 && md <= 4);
	if (md == 2) {
		return cur + ((master_cur - cur) / 2);
	}
	else if (md == 3) {
		return cur + (master_cur - master_prev);
	}
	else if (md == 4) {
		return master_prev;
	}
	return pair<int, int>();
}

vector<string> split(const string& str) {
	istringstream ss(str);
	vector<string> words;
	string word;
	while (ss >> word) {
		words.push_back(word);
	}
	return words;
}

void pp(const string &prefix, const pair<int, int> &p) {
	cout << prefix << p.first << " " << p.second << "\n";
}
