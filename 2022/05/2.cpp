#include <iostream>
#include <string>
#include <vector>
#include <stack>
#include <cassert>

using namespace std;


string extract_argument(const string &exp, const string &command);
void test_func_extract_argument();

const string MOVE_KW = "move";
const string FROM_KW = "from";
const string TO_KW = "to";
const int STACKS_CNT = 9;

int main() {
	test_func_extract_argument();
	string line;
	vector<stack<char>> stacks(STACKS_CNT, stack<char>());
	vector<vector<char>> stacks_v(STACKS_CNT, vector<char>());

	while (getline(cin, line)) {
		int pos;
		if ((pos = line.find('[')) == string::npos) {
			break;
		}

		do {
			stacks_v[(pos >> 2)].push_back(line[pos + 1]);
		} while (((pos = line.find('[', pos + 1)) != string::npos));
	}

	for (int i = 0; i < stacks_v.size(); i++) {
		vector<char> &stack_v = stacks_v[i];
		stack<char> &stack_ = stacks[i];
		for (int j = stack_v.size() - 1; j >= 0; j--) {
			stack_.push(stack_v[j]);
		}
	}
	stacks_v.clear();

	while (getline(cin, line) && line.find(MOVE_KW) == string::npos);
	do {
		int count = stoi(extract_argument(line, MOVE_KW));
		int move_from = stoi(extract_argument(line, FROM_KW)) - 1;
		int move_to = stoi(extract_argument(line, TO_KW)) - 1;

		// cout << "move_from: " << move_from << "move_to: " << move_to << "\n";

		stack<char> &from = stacks[move_from];
		stack<char> &to = stacks[move_to];

		// for (int i = 0; i < count; i++) {
		// 	if (from.empty()) {
		// 		cout << "WTF\n";
		// 	}
		// 	to.push(from.top());
		// 	from.pop();
		// }
		vector<char> v;
		v.reserve(count);
		for (int i = 0; i < count; i++) {
			if (from.empty()) {
				cout << "WTF\n";
			}
			v.push_back(from.top());
			from.pop();
		}
		for (int i = v.size() - 1; i >= 0; i--) {
			to.push(v[i]);
		}

	} while (getline(cin, line));


	string stack_tops;
	for (const stack<char>& s : stacks) {
		if (!s.empty()) {
			stack_tops += s.top();
		}
	}

	cout << stack_tops << "\n";

	return 0;
}

string extract_argument(const string &exp, const string &command) {
	int pos = exp.find(command);
	if (pos == string::npos) {
		return "";
	}

	//assumes only 1 space character between command and its argument
	int arg_pos = pos + command.length() + 1;
	pos = exp.find_first_of(" \n\t\r", arg_pos);
	if (pos == string::npos) {
		pos = exp.length();
	}

	return exp.substr(arg_pos, pos - arg_pos);
}

void test_func_extract_argument() {
	struct inout {
		pair<string, string> in;
		string out;
	};

	vector<inout> t;
	t.push_back({{"command arg", "command"}, "arg"});
	t.push_back({{"command arg ", "command"}, "arg"});
	t.push_back({{"   this this_arg that that_arg other other_arg another another_arg", "other"}, "other_arg"});

	for (const inout &telement : t) {
		assert(extract_argument(telement.in.first, telement.in.second) == telement.out);
	}
}
