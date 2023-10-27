public class Main {
  public static void main(String[] args) {
    try {
      double a = 15.0;
      double b = 0;
      double result = a / b;
      System.out.println(result);
    } catch (ArithmeticException e) {
      System.out.println(e);
      System.out.println("Δεν γίνεται να διαιρέσετε με το μηδέν");
    } finally {
      System.out.println("Finally...");
    }
  }
}
