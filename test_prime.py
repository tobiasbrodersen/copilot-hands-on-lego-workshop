# FILE: /Users/dktobbro/Documents/lego/copilot-hands-on-lego-workshop/test_prime.py

import unittest
from prime import is_prime

class TestIsPrime(unittest.TestCase):
    def test_small_primes(self):
        self.assertTrue(is_prime(2))
        self.assertTrue(is_prime(3))
        self.assertTrue(is_prime(5))
        self.assertTrue(is_prime(7))

    def test_small_non_primes(self):
        self.assertFalse(is_prime(1))
        self.assertFalse(is_prime(4))
        self.assertFalse(is_prime(6))
        self.assertFalse(is_prime(8))
        self.assertFalse(is_prime(9))
        self.assertFalse(is_prime(10))

    def test_large_primes(self):
        self.assertTrue(is_prime(101))
        self.assertTrue(is_prime(103))
        self.assertTrue(is_prime(107))
        self.assertTrue(is_prime(109))

    def test_large_non_primes(self):
        self.assertFalse(is_prime(100))
        self.assertFalse(is_prime(102))
        self.assertFalse(is_prime(104))
        self.assertFalse(is_prime(105))

    def test_edge_cases(self):
        self.assertFalse(is_prime(-1))
        self.assertFalse(is_prime(0))
        self.assertFalse(is_prime(1))

if __name__ == '__main__':
    unittest.main()
