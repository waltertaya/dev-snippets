import random

'''
Description:
	Design a class that stores integers values (with the following methods):
	1. Insert a value (no duplicates)
	2. Removing a value
	3. Get random a value that is already inserted (with equal probability)
'''

class Store:
	def __init__(self):
		self.data = []
		self.data_map = {}

	def insert(self, value):
		if value not in self.data_map:
			self.data.append(value)
			self.data_map[value] = len(self.data) - 1

	def remove(self, value):
		if value in self.data_map:
			idx = self.data_map[value]
			
			if idx != len(self.data) - 1:
				self.data[idx], self.data[-1] = self.data[-1], self.data[idx]
				self.data.pop()
				self.data_map[self.data[idx]] = idx
			else:
				self.data.pop()
			del self.data_map[value]

	def getRandomValue(self):
		if len(self.data) <= 0:
			return 0
		return random.choice(self.data)


if __name__ == '__main__':
	store = Store()
	store.insert(2)
	store.insert(3)
	store.insert(11)
	store.insert(7)
	store.insert(3)
	store.insert(9)
	
	print(store.data)
	store.remove(7)
	print(store.getRandomValue())
	print(store.data)
	print(store.data_map)
	store.remove(9)
	store.remove(99)
	print(store.data)
	print(store.data_map)
