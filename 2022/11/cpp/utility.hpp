#ifndef __UTILITY_HPP__
#define __UTILITY_HPP__
#include <string>
#include <sstream>
#include <vector>
#include <set>
std::vector<std::string> split_by_chars(const std::string &str,
		const std::vector<char> &delims);
std::vector<std::string> split(const std::string&);
#endif
