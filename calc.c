#include <stdio.h>
#include <math.h>

struct {
	int a;
} others;

struct {
    char op;
    float fnum;
    float snum;
    float solu;
} mvar;


void operations() {
    switch (mvar.op) {
        case '+':
            mvar.solu = mvar.fnum + mvar.snum;
            break;
        case '-':
            mvar.solu = mvar.fnum - mvar.snum;
            break;
        case '*':
            mvar.solu = mvar.fnum * mvar.snum;
            break;
        case '/':
            if (mvar.snum == 0) {
                printf("Division through zero is impossible\n");
            } else {
                mvar.solu = mvar.fnum / mvar.snum;
            }
            break;
        case '%':
            mvar.solu = mvar.snum / 100 * mvar.fnum;
            break;
	case '^':
	    mvar.solu = pow(mvar.fnum, mvar.snum);
	case 'v':
	    if (mvar.fnum == 2) {
	    	mvar.solu = sqrt(mvar.fnum);
	    } else {
		mvar.solu = pow(mvar.fnum, 1.0/mvar.snum);
	    	}
	case 'x':
	    others.a = 0;

        default:
            printf("Invalid Input\n");
    }
    printf("%f\n", mvar.solu);
}

int main() {
   others.a = 1;
   char ch;
  while (others.a) { 
       scanf("%f %c %f", &mvar.fnum, &mvar.op, &mvar.snum);
       operations();
    }
    return 0;
}


