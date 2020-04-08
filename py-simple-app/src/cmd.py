import argparse


def parse_args():
    parser = argparse.ArgumentParser(description='Run resource fetcher')
    parser.add_argument('--sum',
                        '-s',
                        dest='port',
                        default=8081,
                        help='Port which apps will listen to')
    args = parser.parse_args()
    return args
