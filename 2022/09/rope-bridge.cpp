#include <string>
#include <iostream>
#include <vector>
#include <set>
#include <sstream>
#include <utility>



using namespace std;

pair<int, int> operator+(const pair<int, int> a, const pair<int, int> b);
pair<int, int> operator-(const pair<int, int> a, const pair<int, int> b);
pair<int, int> get_delta(string d);
int max_abs_reduce(pair<int, int> p);
vector<string> split(const string& str);

void pp(const string &prefix, const pair<int, int> &p); //for debuggin purposes

int main() {
	string line;

	set<pair<int, int>> visited_positions;
	visited_positions.insert(pair<int, int>(0, 0));
	pair<int, int> head_pos = {0, 0}, tail_pos = {0, 0};

	while (getline(cin, line)) {
		auto words = split(line);
		string direction = words[0];
		int moves = stoi(words[1]);

		pair<int, int> delta = get_delta(direction);
		for (int i = 0; i < moves; i++) {
			auto head_pos_prev = head_pos;
			head_pos = head_pos + delta;


			if (max_abs_reduce(head_pos - tail_pos) > 1) {
				tail_pos = head_pos_prev;
			}
			visited_positions.insert(tail_pos);
			// pp("head: ", head_pos);
			// pp("tail: ", tail_pos);
		}
	}

	cout << visited_positions.size() << "\n";
	// for (auto p : visited_positions) {
	// 	pp("", p);
	// }
	return 0;
}

pair<int, int> operator+(const pair<int, int> a, const pair<int, int> b) {
	return pair<int, int>(a.first + b.first, a.second + b.second);
}

pair<int, int> operator-(const pair<int, int> a, const pair<int, int> b) {
	return pair<int, int>(a.first - b.first, a.second - b.second);
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
