
package main
import(
    "fmt"
    "bufio"
    "os"
    "strconv"
)

func main() {
 //Enter your code here. Read input from STDIN. Print output to STDOUT
  scanner := bufio.NewScanner(os.Stdin)
    for i := 1; scanner.Scan(); i++ {
        if i != 1 {
            grade := scanner.Text()
        
            gradeToInt, err := strconv.Atoi(grade)
             if err != nil {
                 fmt.Println(err)
             }
            fmt.Println(evalGrade(gradeToInt))
        }
        
    }
}


func evalGrade(grade int) int {
    if grade < 38 {
        return grade   
    } else {
       x := (grade%5)
	   z := 5

        if (z-x) < 3 {
        grade += (z-x)
      }
    }
    
    return grade
    
}
