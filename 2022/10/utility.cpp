#include "utility.hpp"

using namespace std;

vector<string> split(const string& str) {
	istringstream ss(str);
	vector<string> words;
	string word;
	while (ss >> word) {
		words.push_back(word);
	}
	return words;
}
