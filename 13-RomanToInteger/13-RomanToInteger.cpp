// Last updated: 11/7/2025, 2:49:38 PM
class Solution {
public:
    int romanToInt(string s) {
        char arr[7]={'I','V','X','L','C','D','M'};
        int arr2[7]={1,5,10,50,100,500,1000};
        int num=0;
        for(int i=0;i<s.length();i++){
            for(int j=0;j<7;j++){
                if(s[i]==arr[j]){
                    if((s[i]=='I') || (s[i]=='X') || (s[i]=='C')){
                        if(s[i+1]==arr[j+1]){
                            num=num-arr2[j];
                        }

                        else if(s[i+1]==arr[j+2]){
                            num=num-arr2[j];
                        }

                        else num+=arr2[j];    
                    }
                    else num+=arr2[j];
                
                }
            }
        }
        return num;
    }
};