#include <fstream>
#include <iostream>
#include <string>
int main (int argc, char *argv[]) {
    std::ifstream myfile ("input.txt");
    std::string mystring;
    int red = 0;
    int green = 0;
    int blue = 0;
    int sum = 0;
    if(myfile.is_open()){
        while(true){
            myfile >> mystring;
            if(!myfile.good()){
                break;
            }
            if(mystring=="Game"){
                int multiplier = red*blue*green;
                sum +=multiplier;
                red = 0;
                green = 0;
                blue = 0;
                myfile >> mystring;
            }
            else{
               int number = stoi(mystring);
               myfile >> mystring;
               if(mystring[0] == 'r'){number>red ? red = number: red=red;}
               else if(mystring[0] == 'g'){number>green ? green = number: green=green;}
               else if(mystring[0] == 'b'){number>blue ? blue = number: blue=blue;}
            }
        }
    }
    int multiplier = red*blue*green;
    sum +=multiplier;
    std::cout << sum << std::endl;
    return 0;
}
