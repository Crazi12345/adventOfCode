#include <iostream>
#include <fstream>
#include <iterator>
#include <ostream>
#include <string>
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
                    if(mystring[i]<58 && mystring[i]>47){
                        if(isFirst){
                            first = mystring[i];
                            isFirst = false;
                        }
                        last = mystring[i];
                    }
                    std::cout << mystring[i];
                   }

           char number[2] = {first,last};
           std::cout << "  " << number;
           sum += atoi(number);
           std::cout << std::endl;
        }
    }
    std::cout <<sum << std::endl;
    return 0;

}
