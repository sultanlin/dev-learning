using System;
// help me!
namespace HelloWorld
{
    public class Person
    {
       public string FirstName { get; set; }
       public string LastName { get; set; }

       public string GetFullName()
       {
         return $"{FirstName} {LastName}";
       }
    }
}
