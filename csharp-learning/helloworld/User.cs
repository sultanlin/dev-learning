using System;
using System.Collections.Generic;
using System.Linq;

namespace HelloWorld
{
    public class User
    {

        string firstName = "Sultan";
        public string FirstName { get; set; }

        public string LastName { get; set; }

        public string FullName
        {
            get
            {
                return FirstName + " " + LastName;
            }
        }

        public void Output()
        {
            Console.WriteLine(FirstName);
            Console.WriteLine(LastName);
        }

        public static void PrintUser(User user)
        {
            Console.WriteLine("Static Method. Print User");
            user.Output();
        }

        public static void PrintUsers(List<User> users)
        {
            foreach (User user in users)
            {
                user.Output();
            }
        }
    }
}
