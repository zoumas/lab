abstract class Book {
  protected String ISBN;
  protected String name;
  protected int pages;
  protected double value;

  public static int current_Books = 0;

  public Book(String ISBN, String name, int pages, double value) {
    current_Books++;

    this.ISBN = ISBN;
    this.name = name;
    this.pages = pages;
    this.value = value;
  }

  public int getPages() {
    return pages;
  }

  @Override
  public String toString() {
    return "Book{" +
    "ISBN='" + ISBN + '\'' +
    ", name='" + name + '\'' + 
    ", pages=" + pages +
    ", value=" + value + '}';
  }
}

class JavaBook extends Book {
  private String oracleCD;
  private String author;

  public JavaBook(String ISBN, String name, int pages, double value, String oracleCD, String author) {
    super(ISBN, name, pages, value);
    this.oracleCD = oracleCD;
    this.author = author;
  }

  @Override
  public String toString() {
    return super.toString() +
      "JavaBook{" + 
      "oracleCD='" + oracleCD + '\'' +
      ", author='" + author + "'}";
  }
}

class CppBook extends Book {
  private String versionOfCpp;
  private boolean isPublished;

  public CppBook(String ISBN, String name, int pages, double value, String versionOfCpp, boolean isPublished) {
    super(ISBN, name, pages, value);
    this.versionOfCpp = versionOfCpp;
    this.isPublished = isPublished;
  }

  @Override
  public String toString() {
    return super.toString() + 
    "CppBook{" +
    "versionOfCpp='" + versionOfCpp + '\'' +
    ", isPublished='" + isPublished + "'}";
  }
}

class AllBooks {
  Book[] books;
  private int totalPages;

  public AllBooks(Book[] books) {
    this.books = books;
    calculateTotalPages();
  }
  
  private void calculateTotalPages() {
    totalPages = 0;
    for (Book b : books) {
      totalPages += b.getPages();
    }
  }

  public int getTotalPages() {
    return totalPages;
  }

  public void printAllBooks() {
    for (Book b : books) {
      System.out.println(b);
    }
  }
}

public class Main {
  public static void main(String[] args) {
    JavaBook firstBook = new JavaBook("12345678", "JavaFX", 741, 47, "HerbertCD", "Herbert");
    CppBook secondBook = new CppBook("342356242", "C++17", 1102, 67, "17", true);
    CppBook thirdBook = new CppBook("342356915", "C++20", 1427, 79, "20", false);

    Book[] books = {firstBook, secondBook, thirdBook};

    System.out.println("Συνολικά Βιβλία: " + Book.current_Books);
    AllBooks library = new AllBooks(books);
    System.out.println("Σύνολο Σελίδων: " + library.getTotalPages());
    library.printAllBooks();
  }
}
