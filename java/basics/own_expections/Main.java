public class Main {
  public static void main(String[] args) {
    try {
      int[] a = new int[5];
      int[] primes = {2, 3, 5, 7, 11};
      int[] xs = {1, 2, 3, 5, 7};

      printNumber(primes, 4); // 11

      a = setPrimes(primes);
      System.out.println(a[1]); // 3

      // exceptions
      a = setPrimes(xs);
      printNumber(xs, 6);
    } catch (IndexOutOfBoundsException e) {
      System.out.println(e);
    } catch (myExceptionClass e) {
      System.out.println(e);
    } finally {
      System.out.println("Άπειροι πρώτοι αριθμοί...");
    }
  }

  static void printNumber(int[] a, int index) throws IndexOutOfBoundsException {
    if (index >= a.length) {
      throw new IndexOutOfBoundsException();
    }
    System.out.println("Ο αριθμός είναι: " + a[index]);
  }

  static int[] setPrimes(int[] a) throws myExceptionClass {
    if (a[0] == 1) {
      throw new myExceptionClass("Ο 1 δεν είναι πρώτος αριθμός!");
    }
    return a;
  }
}

class myExceptionClass extends Exception {
  private String message;

  public myExceptionClass(String message) {
    this.message = message;
  }

  @Override
  public String toString() {
    return "myExceptionClass{" + "'message='" + message + "'}";
  }
}
