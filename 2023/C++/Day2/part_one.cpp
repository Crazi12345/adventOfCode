#include <fstream>
#include <iostream>
#include <string>
int main (int argc, char *argv[]) {
    std::ifstream myfile ("input.txt");
    std::string mystring;
    int red = 12;
    int green = 13;
    int blue = 14;
    int sum = 0;
    int iterator = 0;
    int id=0;
    if(myfile.is_open()){
        while(true){
            myfile >> mystring;
            if(!myfile.good()){
                break;
            }
            if(mystring=="Game"){
                iterator++;
                sum+=id;
                id = iterator;
                myfile >> mystring;
            }
            else{
               int number = stoi(mystring);
               myfile >> mystring;
               if(mystring[0] == 'r'){number>red ? id = 0: id=id;}
               else if(mystring[0] == 'g'){number>green ? id = 0: id=id;}
               else if(mystring[0] == 'b'){number>blue ? id = 0: id=id;}
            }
        }
    }
    sum += id;
    return 0;
}
