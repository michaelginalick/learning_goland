require_relative "./cell"
class Game
  
  attr_accessor :number_of_bombs, :grid
  
  def initialize(bombs)
    @number_of_bombs = bombs
    @grid = Array.new(10) { Array.new(10) }
    populate_grid
    seed_bombs
    count_neighbors
    print_grid
    new_game
  end

  def new_game

    game_over = false

    while !game_over do
      input = fetch_input
      game_over = true unless display(input)
      print_grid unless game_over
    end
    reveal_all
    print_grid
  end

  def display(input)
    return false if @grid[input[0].to_i][input[1].to_i].bee == true
    @grid[input[0].to_i][input[1].to_i].reveal(@grid)
    true
  end


  def print_grid
    @grid.each_with_index do |row, i|
      row.each_with_index do |col, j|
        if @grid[i][j].revealed

          if @grid[i][j].bee
            print "B \t"
          elsif @grid[i][j].bee_neighbor_count > 0
            print "#{@grid[i][j].bee_neighbor_count} \t"
          else
            print "#{@grid[i][j].display} \t"
          end
        else
          print "#{@grid[i][j].display} \t"
        end
      end
      puts "\n"
    end
  end

  def fetch_input
    puts "Please enter coordinates"
    input = gets.chomp()
    input = input.split(",")
  end

  private

  def seed_bombs
    while @number_of_bombs > 0
      random_cell = @grid[rand(@grid.length)][rand(@grid[0].length)]
      next if random_cell.bee
      random_cell.bee = true
      @number_of_bombs -=1
    end
  end

  def count_neighbors
    @grid.each_with_index do |row, xOff|
      row.each_with_index do |col, yOff|
        @grid[xOff][yOff].count(@grid)
      end
    end
  end

  def populate_grid
    @grid.each_with_index do |row, i|
      row.each_with_index do |col, j|
        @grid[i][j] = Cell.new(i, j)
      end
    end
    @grid
  end

  def reveal_all
    @grid.each_with_index do |row, i|
      row.each_with_index do |col, j|
        @grid[i][j].revealed = true
      end
    end
  end

end

g = Game.new(10)
