#include "utility.hpp"

using namespace std;

vector<string> split_by_chars(const string &str, const vector<char> &delims) {
	set<char> delims_set(delims.begin(), delims.end());
	vector<string> words;

	auto word_begin = str.begin();
	for (auto it = str.begin(); ; it++) {
		if (it == str.end()) {
			if (it - word_begin > 0)
				words.push_back(str.substr(word_begin - str.begin(), it - word_begin));
			break;
		}
		else {
			const char &c = *it;
			if (delims_set.count(c)) {
				if (it - word_begin > 0) 
					words.push_back(str.substr(word_begin - str.begin(), it - word_begin));
				word_begin = it + 1;
			}
		}
	}
	return words;
}

vector<string> split(const string &str) {
	return split_by_chars(str, {' '});
}
