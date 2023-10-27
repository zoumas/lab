@FunctionalInterface
interface I {
  void hi();
}

class Person {
  String name;
  int age;

  @Deprecated
  public Person(String name) {
    this.name = name;
  }

  public Person(String name, int age) {
    this.name = name;
    this.age = age;
  }

  @Override
  public String toString() {
    StringBuilder b = new StringBuilder();
    b.append("Person{name='");
    b.append(name);
    b.append("', age=");
    b.append(age);
    b.append("}");
    return b.toString();
  }
}

class Man extends Person {
  String style;

  public Man(String name, int age, String style) {
    super(name, age);
    this.style = style;
  }

  @Override
  public String toString() {
    StringBuilder b = new StringBuilder();
    b.append("Person{name='");
    b.append(name);
    b.append("', age=");
    b.append(age);
    b.append("', style='");
    b.append(style);
    b.append("'}");
    return b.toString();
  }
}

public class Main {
  public static void main(String[] args) {
    // Deprecated
    // Person p = new Person("Nick");
    Man m = new Man("Ilias", 21, "Lazy");
    System.out.println(m);
  }
}
