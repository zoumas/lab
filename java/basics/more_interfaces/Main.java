import java.util.Random;

interface I {
  String VERSION = "JDK8";

  static void messageOfDay() {
    System.out.println("Νέες δυνατότητες μετά την " + VERSION);
  }

  default void showRandomNumber() {
    System.out.println("Ο ψευδοτυχαίος αριθμός είναι ο: " + returnRandomNumber());
  }

  private int returnRandomNumber() {
    Random r = new Random();
    return r.nextInt(100) + 1;
  }
}

class C implements I {}

public class Main {
  public static void main(String[] args) {
    I.messageOfDay();
    C c = new C();
    c.showRandomNumber();

  }
}
