import java.io.BufferedReader;
import java.io.FileNotFoundException;
import java.io.FileReader;
import java.io.IOException;

public class Main {
  public static void main(String[] args) {
    if (args.length == 0) {
      return;
    }
    String filename = args[0];

    try (BufferedReader br = new BufferedReader(new FileReader(filename))) {
      while (br.ready()) {
        String line = br.readLine();
        System.out.println(line);
      }
    } catch (IndexOutOfBoundsException | FileNotFoundException e) {
      System.out.println("Πρόβλημα με το αρχείο: " + e);
    } catch (IOException e) {
      System.out.println("Γενικό πρόβλημα: " + e);
    } finally {
      System.out.println("Finally...");
    }
  }
}
