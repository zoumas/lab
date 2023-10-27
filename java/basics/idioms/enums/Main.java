enum Car {
  Mercedes("123", 45_000),
  Porsche("456", 64_000),
  Ferrari("789", 85_000),
  Audi("1011", 45_000),
  Mazda("312", 40_000);

  private String number;
  private double price;

  Car(String number, double price) {
    this.number = number;
    this.price = price;
  }

  public double getPrice() {
    return price;
  }
}

public class Main {
  public static void main(String[] args) {
    Car ferrari = Car.Ferrari;
    System.out.println(Car.valueOf("Ferrari"));
    System.out.println(ferrari);

    try {
      System.out.println(Car.valueOf("Tesla"));
    } catch (IllegalArgumentException e) {
      System.out.println("No such Car: " + e);
    }

    System.out.println("Available Cars");
    for (Car c : Car.values()) {
      System.out.println(c);
    }

    double max = -1;
    Car mostExpensiveCar = null;
    for (Car c : Car.values()) {
      double price = c.getPrice();
      if (price > max) {
        max = price;
        mostExpensiveCar = c;
      }
    }
    System.out.println("The most expensive car is " + mostExpensiveCar + " with price " + max);
  }
}
