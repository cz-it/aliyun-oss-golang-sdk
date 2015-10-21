
.DEFAULT:all

all : 
	cd ossapi; make	
	cd osscmd; make	


.PHONY:clean

clean:
	cd ossapi; make clean;
	cd osscmd; make clean;
