import java.util.ArrayList;
import java.util.Iterator;

class Person {
  private int age;
  private String name;
  private String city;

  public Person(int age, String name, String city) {
    this.age = age;
    this.name = name;
    this.city = city;
  }

  @Override
  public String toString() {
    StringBuilder b = new StringBuilder();
    b.append("Person{age=");
    b.append(age);
    b.append(", name='");
    b.append(name);
    b.append("', city='");
    b.append(city);
    b.append("'}");
    return b.toString();
  }
  
  public int getAge() {
    return age;
  }
}

public class Main {
  public static void main(String[] args) {
    ArrayList<Person> people = new ArrayList<>();
    people.add(new Person(27, "George", "Athens"));
    people.add(new Person(21, "Joan", "Thessaloniki"));
    people.add(new Person(24, "Vasiliki", "Crete"));

    System.out.println("The size of the collection is: " + people.size());
    System.out.println(people);

    int min = people.get(0).getAge();
    for (Person p : people) {
      int age = p.getAge();
      if (age < min) {
        min = age;
      }
    }
    System.out.println("The minimum age is: " + min);

    int max = people.get(0).getAge();
    for (Iterator<Person> iter = people.iterator(); iter.hasNext(); ) {
      int age = iter.next().getAge();
      if (age > max) {
        max = age;
      }
    }
    System.out.println("The maximum age is: " + max);
  }
}
