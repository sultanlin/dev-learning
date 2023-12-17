
using System;
using System.Collections;
using System.Collections.Generic;

namespace leetcodechalls
{
    class Program
    {

        static void Main(string[] args)
        {
            string word1 = "abc";
            string word2 = "pqr123";
            string mergedword = "";
            // (word1.Length + word2.Length)
            Console.WriteLine(word1.Length + word2.Length);
            for (int i = 0; i < (word1.Length + word2.Length); i++)
            {
                if (word1.Length == 0 && word2.Length == 0)
                {
                    break;
                }
                Console.WriteLine(i);
                if (i % 2 == 1 || word1.Length == 0)
                {
                    mergedword += word2[0];
                    Console.WriteLine(mergedword);
                    word2 = word2.Remove(0, 1);
                    Console.WriteLine(word2);
                }
                if (i % 2 == 0 || word2.Length == 0)
                {
                    mergedword += word1[0];
                    Console.WriteLine(mergedword);
                    // word1 = word1.Substring(1);
                    word1 = word1.Remove(0, 1);
                    Console.WriteLine(word1);
                }
            }
            Console.WriteLine(mergedword);
        }
    }
}

// See https://aka.ms/new-console-template for more information
// See https://aka.ms/new-console-template for more information
