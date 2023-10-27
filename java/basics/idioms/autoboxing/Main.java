public class Main {
  public static void main(String[] args) {
    Integer i = 5; // autoboxing with int
    int n = i; // from Integer to int
    n++;
    System.out.println("Primitive: " + n + " Object: " + i);

    Boolean check = true;
    if (check) {
      System.out.println("Autoboxing in Java!");
    }

    Integer num = 254;
    int b = num.byteValue();
    System.out.println(b);

    Double[] grades = {6.0, 9.2, 8.8, 7.4};
    System.out.println(average(grades));
  }

  static Double average(Double[] numbers) {
    Double sum = 0.0;
    for (Double n : numbers) {
      sum += n;
    }
    return sum / numbers.length;
  }
}
