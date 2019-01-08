class Game

  attr_accessor :rolls, :current_role

  def initialize
    @rolls = []
    @current_role = 0
  end


  def roll(pins)
    @rolls[@current_role] = pins
    @current_role +=1
  end

  def score 
    game_score = 0
    frame_index = 0

    while frame_index < 10
    if @rolls[frame_index] == 10
      game_score += 10
      game_score += @rolls[frame_index+1] + @rolls[frame_index+2]
      frame_index +=1
    elsif is_spare(frame_index)
        game_score +=10 + @rolls[frame_index+2]
        frame_index+=2
      else
        game_score += @rolls[frame_index] + @rolls[frame_index+1]
        frame_index+=2
      end
    end

    game_score
  end

  def roll_many(n, pins)

    i = 0

    while i < n
      roll(pins)
      i +=1
    end

  end

  private

  def is_spare(frame_index)
    @rolls[frame_index] + @rolls[frame_index+1] == 10
  end

end

g = Game.new
# g.roll(5)
# g.roll(5)
# g.roll(3)
# g.roll_many(17,0);
# puts g.score() == 16  #spare

g.roll(10)
g.roll(3)
g.roll(4)
g.roll_many(20, 0) #strike

puts g.score
puts g.score == 24 #strike
