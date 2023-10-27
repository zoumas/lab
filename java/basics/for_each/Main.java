public class Main {
  public static void main(String[] args) {
    int[] primes = {5, 11, 13, 17, 23, 19, 2, 7, 3};
    int[][] twoDPrimes = { 
      {5, 11, 13},
      {17, 23, 19},
      {2, 7, 3}
    };

    int min = primes[0];
    for (int x : primes) {
      min = x < min ? x : min;
    }
    System.out.println("Η μικρότερη τιμή του primes: " + min);

    int max = -1;
    for (int xs[] : twoDPrimes) {
      for (int x : xs) {
        if (x > max) {
          max = x;
        }
      }
    }
    System.out.println("Η μεγαλύτερη τιμή του twoDPrimes: " + max);
  }
}
