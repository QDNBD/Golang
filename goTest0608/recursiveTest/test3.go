package main
import "fmt"
//求平方根
//从某个猜测的值 z 开始，我们可以根据 z² 与 x 的近似度来调整 z，产生一个更好的猜测：
//z -= (z*z - x) / (2*z)
func sqrt(x float64,i float64) (float64,float64){
	remain:=(i*i-x)/(2*i);
	i -= remain
	if(remain>0){
		return sqrt(x,i);
	}else{
		return i,remain
	}
}
func get_sqrt(x float64) float64{
	i,_ :=sqrt(x,x);
	return i;
}
func main(){
	fmt.Println(get_sqrt(2))
	fmt.Println(get_sqrt(3))
}