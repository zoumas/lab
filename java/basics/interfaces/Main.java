interface test {
  void welcomeMessage();
  void byeMessage();
  int calculate(int num1, int num2);
}

interface test2 extends test {
  void otherMessage();
}

class firstClass implements test {
  @Override
  public void welcomeMessage() {
    System.out.println("Hello People!");
  }

  @Override
  public void byeMessage() {
    System.out.println("Bye People!");
  }

  @Override
  public int calculate(int num1, int num2) {
    return num1 % num2;
  }
}

class secondClass extends firstClass implements test2 {
  @Override
  public void otherMessage() {
    System.out.println("Java is powerful!");
  }
}

public class Main {
  public static void main(String[] args) {
    firstClass fc = new firstClass();
    System.out.println(fc.calculate(7, 5));
    fc.byeMessage();
    secondClass sc = new secondClass();
    sc.otherMessage();

  }
}
