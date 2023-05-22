create or replace function pg_dijkstra_go(text,int,int) 
  returns text as :MOD, 'pgDijkstraGo' 
  language c strict;
