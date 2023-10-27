public class Main {
  public static void main(String[] args) {
    // Ενά πολύ κακό παράδειγμα
    if (0xFFFF_FFFF != -1) {
      System.out.println("1");
    } else if ((0xFFFF_FFFF >>> 24) == 255) {
      if ((0xFFFF_FFFF & 10) == 10) {
        System.out.println("2");
      }
    } else if ((0xFFFF_FFFF >> 10) == -1){
      System.out.println("3");
    } else {
      System.out.println("4");
    }

    String name = "Anastasia";

    if (name.equals("Anna")) {
      System.out.println("Hello Anna");
    } else if (name.equals("Maria")) {
      System.out.println("Hello Maria");
    } else if (name.equals("Anastasia")) {
      System.out.println("Hello Anastasia!");
    }
  }
}
