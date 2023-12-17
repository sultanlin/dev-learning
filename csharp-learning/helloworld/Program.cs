using System;
using System.Collections;
using System.Collections.Generic;

namespace HelloWorld
{
    class Program
    {

        static void Main(string[] args)
        {
            string[] dogs = { "sultan", "jafar", "kalb", "haiwan", };

            var dogSpaces = from dog in dogs
                            where dog.Contains('n')
                            orderby dog descending
                            select dog;

            foreach (string i in dogSpaces)
            {
                Console.WriteLine(i);
            }
        }
    }
}

// See https://aka.ms/new-console-template for more information
