import random

def random_generator(n=100):
    """ generates `n` random values between 0 - 100 (inclusive).
    """
    print "Generator invoked (n==%d)" % (n)
    for i in range(n):
        yield random.randint(0, 100)
