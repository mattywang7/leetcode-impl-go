class MyCalendarTwo:
    def __init__(self) -> None:
        self.booked = []
        self.overlaps = []

    def book(self, start: int, end: int) -> bool:
        if any(s < end and start < e for s, e in self.overlaps):
            return False
        for s, e in self.booked:
            if s < end and start < e:
                self.overlaps.append((max(s, start), min(e, end)))
        self.booked.append((start, end))
        return True


if __name__ == '__main__':
    s = MyCalendarTwo()
    print(s.book(10, 20))
    print(s.book(50, 60))
    print(s.book(10, 40))
    print(s.book(5, 15))
    print(s.book(5, 10))
    print(s.book(25, 55))
