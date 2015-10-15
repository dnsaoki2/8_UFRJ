from random import randint
from math import ceil

GLOBAL_START = 0
GLOBAL_END = 24 * 3600 #seconds
MAX_DURATION = 5 * 3600 #seconds
MIN_DURATION = 60	

def generate_tuples(n):
	tuples = []
	for i in range(n):
		duration = randint(MIN_DURATION, MAX_DURATION)
		start = randint(GLOBAL_START - duration, GLOBAL_END)
		end = start + duration
		tuples += [(start,end)]
	return tuples

def compute_histogram_naive(player_times, columns):
	step = (GLOBAL_END - GLOBAL_START) // columns 
	histogram = []
	moment = GLOBAL_START
	while moment <= GLOBAL_END: 
		player_online = 0
		for times_tuple in player_times:
			if times_tuple[0] <= moment and times_tuple[1] >= moment:
				player_online += 1
		histogram += [(moment, player_online)]
		moment += step
	return histogram

def compute_histogram_linear(player_times, columns):
	step = (GLOBAL_END - GLOBAL_START) // columns 
	histogram = []
	players_by_start_time = {}
	players_by_end_time = {}
	for times_tuple in player_times:
		start = times_tuple[0]
		end = times_tuple[1]
		if end < GLOBAL_START:
			continue
		start = max(start, GLOBAL_START)
		adjusted_start = GLOBAL_START + step * \
											ceil((start - GLOBAL_START) / step)
		adjusted_end = GLOBAL_START + step * \
											ceil((end - GLOBAL_START) / step)
		players_by_start_time[adjusted_start] = \
				players_by_start_time.get(adjusted_start, 0) + 1
		players_by_end_time[adjusted_end] = \
				players_by_end_time.get(adjusted_end, 0) + 1								
	moment = GLOBAL_START
	players_online = 0
	while moment <= GLOBAL_END:
		players_online += players_by_start_time.get(moment,0)
		players_online -= players_by_end_time.get(moment,0)
		histogram += [(moment, players_online)]
		moment += step
	return histogram

def print_histogram(histogram):
	for entry in histogram:
		print("time = %d, players = %d" % (entry[0],  entry[1]))

def main():
	player_times = generate_tuples(10)
	histogram = compute_histogram_naive(player_times, 24)
	print_histogram(histogram)

if __name__ == "__main__":
	main()