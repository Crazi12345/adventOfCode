#include <iostream>
#include <fstream>
#include <iterator>
#include <ostream>
#include <string>

std::string custom_substring(int start, int end, std::string full_string){
    if(end>full_string.size()){
            return "";
        }
    std::string result; 
    int j = 0;
    for(int i = start; i < end; i++){
        result+= full_string[i]; 
        j++;
    }
    return result;
}
std::string custom_replace(int start, int end, std::string full_string, char newChar){
    if(end>full_string.size()){
        std::cout << "TOO BIG" << std::endl;
        return full_string;
    }
   std::string result;
    result = custom_substring(0,start,full_string);

    result+=newChar;
    result+= custom_substring(end,full_string.size(),full_string);
    return result;

}
int string_to_int(std::string number){
    if(number == "one"){
        return 1;
    }
    else if(number == "two"){
        return 2;
    }
    else if(number == "three"){
        return 3;
    }
    else if(number == "four"){
        return 4;
    }
    else if(number == "five"){
        return 5;
    }
    else if(number == "six"){
        return 6;
    }
    else if(number == "seven"){
        return 7;
    }
    else if(number == "eight"){
        return 8;
    }
    else if(number == "nine"){
        return 9;
    }
    return -1;
}
int main(){

    std::ifstream myfile ("input.txt");
    std::string mystring;
    int sum = 0;
    char first = ' ';
    char last = ' ';
    bool isFirst = true;
    if(myfile.is_open()){
        while(true){
            myfile >> mystring;
            isFirst = true;
            if(!myfile.good()){
                break;
            }
            for(int i = 0; i<mystring.size();i++){
                for(int j = 3; j<6;j++){
                    int num = string_to_int(mystring.substr(i,j));
                    if(num != -1){
                        mystring.insert(i,std::to_string(num));
                        i++;
                    }
                }
            }
           for(int i = 0; i<mystring.size();i++){
                    if(mystring[i]<58 && mystring[i]>47){
                        if(isFirst){
                            first = mystring[i];
                            isFirst = false;
                        }
                        last = mystring[i];
                    }
                   }
 char first = '0', last = '0';
            for (char c : mystring) {
                if (isdigit(c)) {
                    if (first == '0') first = c;
                    last = c;
                }
            }

            std::string numberStr = {first, last};
            int number = std::stoi(numberStr);  
            std::cout << first << last << " " << mystring << "  " << (number)<<std::endl;
           sum +=(number);
        }
    }
    std::cout <<sum << std::endl;
    return 0;

}

