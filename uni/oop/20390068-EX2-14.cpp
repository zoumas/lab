/* Άσκηση 2
 *
 * Κλάσεις 
 * (χαρακτηριστικά, μέθοδοι, τροποποιητές πρόσβασης, κατασκευαστές, καταστροφείς)
 *
 * Ονοματεπώνυμο: ΗΛΙΑΣ ΖΟΥΜΑΣ
 * Αριθμός Μητρώου: 20390068
 * Τμήμα: 14
*/

#include <iostream>
#include <string>
#include <cstring>

#include <stdexcept>
#include <cassert>

using std::cout;
using std::cerr;
using std::endl;
using std::string;
using std::ostream;

/// Κλάση Φοιτητής
class Student {
private:
    /// Αριθμός Μητρώου
    char* id;
    /// Ονοματεπώνυμο
    string name;
    /// Τρέχον Εξάμηνο
    unsigned current_semester;
public:
    Student(const char *id, string name, unsigned current_semester = 1);
    Student(const Student &s);
    Student& operator=(const Student &s);
    ~Student();

    char* get_id() const;
    void set_id(const char*);
    string get_name() const;
    void set_name(string);
    unsigned int get_current_semester() const;
    void set_current_semester(unsigned);

    Student& operator++();
    Student operator++(int);

    Student& operator+=(unsigned);
    Student& operator-=(unsigned);
};

/// Εκτύπωση όλων των χαρακτηριστικών, σε μία γραμμή, σε οποιοδήποτε κανάλι (οstream)
ostream& operator<<(ostream &os, const Student &s)
{
    os << s.get_id() << ", "
       << s.get_name() << ", "
       << s.get_current_semester();

    return os;
}

Student::Student(const char *id, string name, unsigned current_semester):
    name(name),
    current_semester(current_semester)
{
    if (id == NULL) {
        throw std::invalid_argument("in Student Constructor: Field id cannot be NULL");
    }
    this->id = new char[strlen(id) + 1];
    strcpy(this->id, id);
}

Student::Student(const Student &s):
    name(s.name),
    current_semester(s.current_semester)
{
    id = new char[strlen(s.id) + 1];
    strcpy(id, s.id);
}

Student& Student::operator=(const Student &s)
{
    delete [] id;

    id = new char[strlen(s.id) + 1];
    strcpy(id, s.id);

    name = s.name;
    current_semester = s.current_semester;

    return *this;
}

Student::~Student()
{
    delete [] id;
}

char* Student::get_id() const { return id; }
void Student::set_id(const char *id)
{
    if (id == NULL) {
        throw std::invalid_argument("in Student set_id: Field id cannot be NULL");
    }

    delete [] this->id;
    this->id = new char[strlen(id) + 1];
    strcpy(this->id, id);
}

string Student::get_name() const { return name; }
void Student::set_name(string n) { name = n; }

unsigned Student::get_current_semester() const { return current_semester; }
void Student::set_current_semester(unsigned cs) { current_semester = cs; }

/// Αύξηση του τρέχοντος εξαμήνου κατά 1 με χρήση του τελεστή προαύξησης ή του τελεστή μετααύξησης.
// prefix 
Student& Student::operator++()
{
    ++current_semester;
    return *this;
}

// postfix
Student Student::operator++(int)
{
    Student tmp = *this;
    current_semester++;
    return tmp;
}

Student& Student::operator+=(unsigned incr)
{
    current_semester += incr;
    return *this;
}
Student& Student::operator-=(unsigned decr)
{
    if (( (int)current_semester - (int)decr ) <= 1) {
        current_semester = 1;
    } else {
        current_semester -= decr;
    }
    return *this;
}

void basic_demonstration();
void test_bad_inputs();

int main(void)
{
    basic_demonstration();

    // test_bad_inputs();
}

bool is_student_equal(const Student s1, const Student s2)
{
    bool res;
    res = (strcmp(s1.get_id(), s2.get_id()) == 0) ? true : false;
    res = s1.get_name() == s2.get_name();
    res = s1.get_current_semester() == s2.get_current_semester();
    return res;
}

// Συνάρτηση που επιδεικνύει τη λειτουργικότητα της Κλάσης Student
void basic_demonstration()
{
    // Assert that the Default value for current_semester is 1
    Student foo("0", "foo");
    assert(foo.get_current_semester() == 1);

    // Constructor with all attributes
    Student hz("20390068", "Ηλίας Ζουμάς", 6);
    cout << hz << endl;

    // Copy Constructor
    Student bar = foo;
    assert(is_student_equal(bar, foo));

    // Assignment Operator
    foo = bar = hz;
    assert(is_student_equal(hz, foo));
    assert(is_student_equal(hz, bar));
    assert(is_student_equal(foo, bar));


    // Assertions για την επαλήθευση της σωστής
    // λειτουργίας των τελεστών της προαύξησης και μετααύξησης.
    assert(hz.get_current_semester() == 6);
    assert((++hz).get_current_semester() == 7);
    assert((hz++).get_current_semester() == 7);
    assert(hz.get_current_semester() == 8);

    // Assertions για την επαλήθευση της σωστής
    // λειτουργίας των τελεστών += και -=.
    assert((hz-=4).get_current_semester() == 4);
    assert((hz+=2).get_current_semester() == 6);
}

void test_bad_inputs()
{
    /*
    // Student foo(NULL, "NULL"); // Student::Student throws an exception
    // Student foo("bad name", NULL); std::string throws it's own exception
    */

    /*
    Student foo ("0", "foo");
    // foo.set_id(NULL);
    foo.set_current_semester(-1); // bad value that the compiler should statically catch but instead wraps around
    cerr << foo << endl;
    */

    /*
    Student hz ("20390068", "Ηλίας Ζουμάς", 6);
    hz-=10;
    cerr << hz << endl;
    */
}
