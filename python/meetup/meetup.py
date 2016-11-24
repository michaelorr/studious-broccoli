import calendar
import datetime


def meetup_day(year, month, day_of_week, count):
    weekdays = {
        0: 'Monday',
        1: 'Tuesday',
        2: 'Wednesday',
        3: 'Thursday',
        4: 'Friday',
        5: 'Saturday',
        6: 'Sunday',
    }
    for i in xrange(31):
        try:
            d = datetime.date(year=year, month=month, day=i + 1)
        except ValueError:
            continue
        if weekdays[d.weekday()] == day_of_week:
            if count == 'teenth':
                if d.day >= 13 and d.day <= 19:
                    return d
            if count == '1st':
                if d.day >= 1 and d.day <= 7:
                    return d
            if count == '2nd':
                if d.day >= 8 and d.day <= 14:
                    return d
            if count == '3rd':
                if d.day >= 15 and d.day <= 21:
                    return d
            if count == '4th':
                if d.day >= 22 and d.day <= 28:
                    return d
            if count == '5th':
                if d.day >= 29:
                    return d
            if count == 'last':
                num_days = calendar.monthrange(year, month)[1]
                if d.day > num_days - 7 and d.day <= num_days:
                    return d
    raise MeetupDayException()


class MeetupDayException(Exception):
    pass
